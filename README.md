# go-inventory-mangement-system

**Project Idea: Inventory Management System**

![image](https://github.com/JustJordanT/go-inventory-mangement-system/assets/38886930/8af4f3ef-c2e4-4cae-85fc-decaac5a7cdf)


**Description:**
Create an inventory management system for a small store using Go. This system will allow you to add, update, delete, and view products in the inventory. You'll use structs to represent products and maps to store and manage the inventory.

**Key Features:**

1. **Product Struct:** Create a `Product` struct with fields like `ID`, `Name`, `Price`, `Quantity`, and any other relevant information for each product.

2. **Inventory Map:** Use a map to store the inventory. The map key could be the product's `ID`, and the value would be the `Product` struct.

3. **Add Products:** Implement a function to add new products to the inventory. Users should be able to input the product details, and a new `Product` struct should be added to the map.

4. **Update Products:** Allow users to update existing products in the inventory. They should be able to search for a product by its `ID` and update its information, such as price or quantity.

5. **Delete Products:** Implement a function to delete products from the inventory. Users can provide the `ID` of the product to be removed.

6. **View Inventory:** Create a function to display the entire inventory, listing all the products and their details.

7. **Search Products:** Implement a search function that allows users to find products by name, price range, or any other relevant criteria.

8. **Save and Load:** You can add functionality to save the inventory to a file (e.g., JSON or CSV) and load it back when the program starts.

**Additional Features (Optional):**

9. **Inventory Reports:** Generate reports, such as the total value of the inventory, the most expensive product, or products that are running low in stock.

10. **User Interface:** If you want to make the project more user-friendly, consider building a simple command-line or web-based user interface for interacting with the inventory system.

## Stack

- [Go](Go.dev)
- [Cobra](https://github.com/spf13/cobra)
- [Charm](https://github.com/charmbracelet)
