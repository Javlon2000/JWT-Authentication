# JWT-Authentication
This is the JWT Authentication program, used PostgreSQL database, Gin web framework, GORM library, WebSocket. 
Steps you need to do:

1) Sign-up:
  POST your email, if your email exits in my database, you will get 400-Bad Request. 
  If not, then you will get 200-OK and also I will send you password in UUID form.

2) Login:
  POST your email and password that I sent you. If your email and password will not be wrong, you will get JSON Web Token.

3) Change Password:
  POST your email, your current password the password that I sent your email, and your new password.
  If everything is fine, you will get new JSON Web Token.
  
4) Echo:
  Verify with new JSON Web Token, using "accessToken" header. 
