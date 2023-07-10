package http

func (s *Server) SetupRoutes() {
	v1 := s.App.Group("/api/v1")

	v1.POST("/sign-up", s.handler.User.SignUp)                              // anyone  Registration
	v1.POST("/sign-in", s.handler.User.SignIn, s.auth.SignInMiddleware)     // anyone  Login
	v1.POST("/update-password", s.handler.User.UpdatePassword, s.auth.Auth) // user  Update Password

	v1.POST("/products", s.handler.Product.CreateProduct, s.auth.Auth)       // seller  Create Product
	v1.GET("/products", s.handler.Product.ListProducts)                      // anyone  List All Product
	v1.GET("/products/:id", s.handler.Product.GetProduct)                    // anyone  Get Product by ID
	v1.PUT("/products/:id", s.handler.Product.UpdateProduct, s.auth.Auth)    // seller  Update Product by ID
	v1.DELETE("/products/:id", s.handler.Product.DeleteProduct, s.auth.Auth) // seller  Delete Product by ID

	v1.GET("/categories", s.handler.Category.ListCategories)  // anyone  List All Categories
	v1.GET("/categories/:id", s.handler.Category.GetCategory) // anyone  Get Category by ID

	v1.POST("/products/:id/reviews", s.handler.Review.CreateReview, s.auth.Auth) // user  Create Review for Product
	v1.GET("/products/:id/reviews", s.handler.Review.ListReviews)                // anyone  List Reviews of one product
	v1.GET("/reviews/:id", s.handler.Review.GetReview)                           // anyone  Get Review by ID
	v1.PUT("/reviews/:id", s.handler.Review.UpdateReview, s.auth.Auth)           // user  Update Review by ID
	v1.DELETE("/reviews/:id", s.handler.Review.DeleteReview, s.auth.Auth)        // user  Delete Review by ID

	v1.POST("/order", s.handler.Order.MakeOrder, s.auth.Auth) // user Make a order
	v1.GET("/order", s.handler.Order.ListOrders, s.auth.Auth) // user  Get user's orders
	v1.GET("/order/:id", s.handler.Order.GetOrder)            // anyone Get order by ID

	v1.POST("/cart", s.handler.Cart.AddToCart, s.auth.Auth)            // user  Add product to user's cart
	v1.GET("/cart", s.handler.Cart.GetUsersCart, s.auth.Auth)          // user  Get products in user's cart
	v1.DELETE("/cart/:id", s.handler.Cart.DeleteFromCart, s.auth.Auth) // user  Delete product from user's cart

	// admin routes
	v1.POST("/admin/categories", s.handler.Category.CreateCategory, s.auth.Auth)       // admin
	v1.PUT("/admin/categories/:id", s.handler.Category.UpdateCategory, s.auth.Auth)    // admin
	v1.DELETE("/admin/categories/:id", s.handler.Category.DeleteCategory, s.auth.Auth) // admin

	v1.PUT("/admin/order/:id", s.handler.Order.UpdateOrder, s.auth.Auth)    // admin
	v1.DELETE("/admin/order/:id", s.handler.Order.DeleteOrder, s.auth.Auth) // admin
}
