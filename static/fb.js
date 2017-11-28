window.fbAsyncInit = function() {
  FB.init({
    appId : '367804320339821',
    cookie : true,
    xfbml : true,
    version : 'v2.11'
  });
  FB.AppEvents.logPageView();
};

(function(d, s, id) {
  var js, fjs = d.getElementsByTagName(s)[0];
  if (d.getElementById(id)) {
    return;
  }
  js = d.createElement(s);
  js.id = id;
  js.src = "https://connect.facebook.net/en_US/sdk.js";
  fjs.parentNode.insertBefore(js, fjs);
}(document, 'script', 'facebook-jssdk'));

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
