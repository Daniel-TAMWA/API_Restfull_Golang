
Our API is a REST API so the functions are CRUD including create, read, update and delete a customer from our slice.
In order to execute our main.go file, you should already be in the folder where it is located by entering the commands 
cd Goproject
cd CRM_backend
go run main.go

In order to test the API, we could use a tool such as Postman or use the Curl command. In our case we used Postman because of its graphical interface. To do this after choosing the method enter the address :
GET http://localhost:3000/customers to list the customers 
Example getCustomers.png in images's folder

GET http://localhost:3000/customers/{id} to select a customer
Example get_one_customer.png in images's folder

POST http://localhost:3000/customers to add a customer
Example addCustomer.png in images's folder

DELETE http://localhost:3000/customers/{id} to delete a customer
Example deleteCustomer.png in images's folder

PUT http://localhost:3000/customers/{id} to update a customer
Example updateCustomer.png in images's folder

Concerning the packages, to install them I used the command go get <package name>

The API is structured as follows:
we have defined a structure named Customer with the fields Id, Name, Role, Phone and Conctated. We made sure that Id is unique using the google uuid package then we created the different functions getCustomers, getCustomer, addCustomer, deleteCustomer and updateCustomer then a function for the different endpoints using the mux package. Finally we have the main function in which we have initialized the values of Customers finally still in the main function call the handle function that manages the endpoints.

