## Firebase Authentication:

Most apps need to know the identity of a user. Knowing a user's identity allows an app to securely save user data in the cloud and provide 
the same personalized experience across all of the user's devices.

Firebase Authentication provides backend services, easy-to-use SDKs, and ready-made UI libraries to authenticate users to your app. 
It supports authentication using passwords, phone numbers, popular federated identity providers like Google, Facebook and Twitter, and more.

## How does it work?

o sign a user into your app, you first get authentication credentials from the user. These credentials can be the user's email address and password, or an OAuth token from a federated identity provider. Then, you pass these credentials to the Firebase Authentication SDK. Our backend services will then verify those credentials and return a response to the client.

After a successful sign in, you can access the user's basic profile information, and you can control the user's access to data stored in other Firebase products. You can also use the provided authentication token to verify the identity of users in your own backend services.
