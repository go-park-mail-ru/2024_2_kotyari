package product

import (
	"fmt"
	"github.com/go-park-mail-ru/2024_2_kotyari/internal/model"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (pd *ProductsDelivery) getAddProductForm(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	formHTML := `
        <html lang="ru">
        <body>
        <form action="/add_product" method="post" enctype="multipart/form-data"  accept-charset="UTF-8">
            Seller ID: <input type="text" name="seller_id"><br>
            Title: <input type="text" name="title"><br>
            Description: <input type="text" name="description"><br>
            Price: <input type="number" name="price"><br>
            Count: <input type="number" name="count"><br>
            Original Price: <input type="number" name="original_price"><br>
            Discount: <input type="number" name="discount"><br>
            Characteristics: <input type="text" name="characteristics" placeholder="e.g. Color=Red, Size=L"><br>
            Categories (Comma Separated IDs): <input type="text" name="categories"><br>
            Images: <input type="file" name="images" multiple><br>
            Options: <br>
            <div id="options-container">
                <div class="option-block">
                    Title: <input type="text" name="option_title[]"><br>
                    Type: <input type="text" name="option_type[]"><br>
                    Option Values (Comma Separated): <input type="text" name="option_values[]"><br>
                </div>
            </div>
            <button type="button" onclick="addOptionBlock()">Add Option</button><br>
            <input type="submit" value="Add ProductBase">
        </form>
        <script>
            function addOptionBlock() {
                let container = document.getElementById('options-container');
                let newBlock = document.createElement('div');
                newBlock.classList.add('option-block');
                newBlock.innerHTML = 'Title: <input type="text" name="option_title[]"><br>Type: <input type="text" name="option_type[]"><br>Option Values (Comma Separated): <input type="text" name="option_values[]"><br>';
                container.appendChild(newBlock);
            }
        </script>
        </body>
        </html>`
	_, err := w.Write([]byte(formHTML))
	if err != nil {
		return
	}
}

func (pd *ProductsDelivery) AddProducts(w http.ResponseWriter, r *http.Request) {
	pd.log.Info("1213")

	if r.Method == http.MethodPost {
		fmt.Println("1234543")
		pd.addProduct(w, r)
	}
	if r.Method == http.MethodGet {
		pd.getAddProductForm(w, r)
	}
}

func (pd *ProductsDelivery) addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // Max upload of 10 MB
	if err != nil {
		pd.log.Error("[ ProductHandler.AddProduct ] Error parsing form", slog.String("error", err.Error()))
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	var card model.ProductCard
	card.Title = r.FormValue("title")
	card.Description = r.FormValue("description")
	card.Count = parseUintOrZero(r.FormValue("count"))
	card.OriginalPrice = parseUintOrZero(r.FormValue("original_price"))
	card.Discount = parseUintOrZero(r.FormValue("discount"))

	card.Price = uint32(float64(card.OriginalPrice) * (1 - float64(card.Discount)/100))
	card.Seller.ID = parseUintOrZero(r.FormValue("seller_id"))

	// Process and save images
	images := r.MultipartForm.File["images"]
	for _, imgHeader := range images {
		// Open the uploaded file
		file, err := imgHeader.Open()
		if err != nil {
			pd.log.Error("[ ProductHandler.AddProduct ] Error opening image file", slog.String("error", err.Error()))
			http.Error(w, "Error opening image file", http.StatusBadRequest)
			return
		}

		// Create a temporary file for image processing
		tempFile, err := os.CreateTemp("", "upload-*"+filepath.Ext(imgHeader.Filename))
		if err != nil {
			pd.log.Error("[ ProductHandler.AddProduct ] Error creating temp file", slog.String("error", err.Error()))
			http.Error(w, "Error creating temp file", http.StatusInternalServerError)
			file.Close() // Close file before returning to avoid leaks
			return
		}

		// Copy content from the uploaded file to the temporary file
		_, err = io.Copy(tempFile, file)
		file.Close() // Close uploaded file after copying content
		if err != nil {
			pd.log.Error("[ ProductHandler.AddProduct ] Error copying file content", slog.String("error", err.Error()))
			http.Error(w, "Error copying file content", http.StatusInternalServerError)
			tempFile.Close()           // Close tempFile before removing
			os.Remove(tempFile.Name()) // Remove temp file before returning
			return
		}

		// Seek to the beginning of the tempFile for reading
		_, err = tempFile.Seek(0, 0)
		if err != nil {
			pd.log.Error("[ ProductHandler.AddProduct ] Error seeking temp file", slog.String("error", err.Error()))
			http.Error(w, "Error seeking temp file", http.StatusInternalServerError)
			tempFile.Close()
			os.Remove(tempFile.Name())
			return
		}

		// Save the image and retrieve its unique name
		imageName, err := pd.imagesUsecase.SaveImage(imgHeader.Filename, tempFile)
		tempFile.Close()           // Close temp file after processing
		os.Remove(tempFile.Name()) // Remove temp file after processing
		if err != nil {
			pd.log.Error("[ ProductHandler.AddProduct ] Error saving image file", slog.String("error", err.Error()))
			http.Error(w, "Error saving image file", http.StatusInternalServerError)
			return
		}

		// Add image name to product card
		card.Images = append(card.Images, model.Image{Url: imageName})
	}

	// Additional processing for characteristics, categories, and options
	// Handling categories
	categoryIDs := strings.Split(r.FormValue("categories"), ",")
	for _, idStr := range categoryIDs {
		id := parseUintOrZero(idStr)
		card.Categories = append(card.Categories, model.Category{ID: id})
	}

	// Handling options
	optionTitles := r.Form["option_title[]"]
	optionTypes := r.Form["option_type[]"]
	optionValuesRaw := r.Form["option_values[]"]

	for i := range optionTitles {
		optionBlock := model.OptionsBlock{
			Title: optionTitles[i],
			Type:  optionTypes[i],
		}
		values := strings.Split(optionValuesRaw[i], ",")
		for _, value := range values {
			optionBlock.Options = append(optionBlock.Options, model.Option{Value: strings.TrimSpace(value)})
		}
		card.Options.Values = append(card.Options.Values, optionBlock)
	}
	card1 := card
	card1.Images = nil
	pd.log.Debug("Полученная карточка",
		slog.Any("card", card1),
	)

	// Save the product card in the database
	err = pd.repo.AddProduct(r.Context(), card)
	if err != nil {
		pd.log.Error("[ ProductHandler.AddProduct ] Error adding product to database", slog.String("error", err.Error()))
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Product added successfully"))
	if err != nil {
		pd.log.Error("[ ProductHandler.AddProduct ] Error writing response", slog.String("error", err.Error()))
		return
	}
}

func parseUintOrZero(val string) uint32 {
	parsed, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return 0
	}
	return uint32(parsed)
}
