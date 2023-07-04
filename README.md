# eCommerce-smart_gads
Smart Gads
==========
Smart Gads is an eCommerce web application built using Go Gin, PostgreSQL, and RESTful API architecture. It aims to provide an easy and enjoyable shopping experience for customers, as well as a convenient way for merchants to manage their products and orders.

Features
========
1. User authentication: Users can create an account, log in, and update their profile information. Passwords are hashed and salted for security.
2. Product management: Merchants can add, edit, and delete their products, including images, descriptions, prices, and categories. Customers can browse and search for products, filter by category and price range, and add them to their cart.
3. Order management: Merchants can view and fulfill orders, mark them as shipped, and update their status. Customers can view their order history and track their shipments.
4. Payment integration: Customers can securely checkout using Stripe, a popular payment processing service. Orders are automatically created and updated with payment information.
5. Admin panel: Administrators can manage users, products, orders, and categories. They can also view statistics and generate reports.
6. API documentation: The RESTful API is well-documented using Swagger, with detailed explanations and examples of each endpoint.
7. Testing and deployment: The codebase is thoroughly tested using Go's testing framework, and automated with continuous integration and deployment tools like GitHub Actions and Docker.
8. Makefile: The project includes a Makefile that allows you to easily build and run the application with the make run command. Other commands available include make test for running tests and make clean for cleaning up build artifacts.


Technologies
============
> Go Gin: A lightweight and fast web framework for Go.
> PostgreSQL: A powerful and open-source relational database.
> RESTful API: A standard architecture for building web APIs.
> Swagger: A tool for designing, building, and documenting APIs.
> GitHub Actions: A tool for automating workflows and continuous integration/deployment.


Getting started
===============
To run the Smart Gads web app locally using the Makefile, you can follow these steps:

1. Clone the repository and navigate to the project directory.
2. Set up the database schema by running psql -f db/schema.sql.
3. Set the required environment variables by creating a .env file (see .env.example for an example).
4. Start the server by running make run.

You can then access the web app at http://localhost:3000 and the Swagger documentation at http://localhost:3000/swagger/index.html.
