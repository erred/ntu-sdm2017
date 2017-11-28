function logout() {
  FB.logout(function(response) {
    window.location.replace("/");
    // TODO revoke session cookie
    // TODO Notify serve of logout
  });
}

function statusChangeCallback(response) {
  switch (response.status) {
  case 'connected':
    console.log('connected!');
    console.log(response);
    // response.authResponse.accessToken;
    // response.authResponse.userID;
    // TODO change login to 'continue to app'
    // TODO send id to server, set session cookie
    break;
  case 'not_authorized':
  case 'unkown':
  default:
    console.log('not connected');
    console.log(response);
    // TODO reset login button and spinner
    break;
  }
}

function checkLoginState() {
  FB.getLoginStatus(function(response) { statusChangeCallback(response); });
}

FB.getLoginStatus(function(response) { statusChangeCallback(response); });
