package main

// GET /api/products - Retrieves a list of all products available in the store, including their names, prices, descriptions, and images.

// GET /api/products/{id} - Retrieves a specific product by its ID, including its name, price, description, and image.

// POST /api/orders - Creates a new order for a customer, including the customer's name, email address, shipping address, payment amount, and payment currency. This endpoint should also generate a Kaspi Pay checkout URL for the customer to complete the payment.

// GET /api/orders - Retrieves a list of all orders that have been placed in the store, including the customer's name and email address, the order total, and the payment status (e.g. pending, completed, refunded).

// GET /api/orders/{id} - Retrieves a specific order by its ID, including the customer's name and email address, the order total, and the payment status.

// PUT /api/orders/{id}/status - Updates the payment status of a specific order, marking it as either completed or refunded
