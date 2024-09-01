// Challenge 11 (with MySQL): User Authentication System
// 1. User Registration and Management:
//    a. Create a users table with fields: id, username, and email.
//    b. Implement POST /users to create a new user.
//    c. Implement GET /users/{id} to fetch a user's details, including a list of their posts.
// 2. Blog Posts:
//    a. Create a posts table with fields: id, user_id (foreign key from users), title, and content.
//    b. Implement POST /posts to create a new post by a user.
//    c. Implement GET /posts/{id} to fetch a post’s details, including the author and comments.
// 3. Comments:
//    a. Create a comments table with fields: id, post_id (foreign key from posts), user_id (foreign key from users), and content.
//    b. Implement POST /comments to add a comment to a post.
//    c. Implement GET /posts/{id}/comments to fetch all comments for a post, including the commenter's username.
// 4. Transactions:
//    Ensure that creating a post and its associated comments is done within a transaction, so either all succeed, or none do.
// 5. Advanced Query:
//    Implement a query that retrieves all posts made by a specific user, along with the number of comments on each post.  Create an endpoint that requires a valid JWT token to access, representing a protected resource.

package main

func main() {

}
