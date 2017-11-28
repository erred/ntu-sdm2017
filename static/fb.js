// ================== StatusChange Handler ======================
function statusChangeCallback(response) {
  switch (response.status) {
  case 'connected':
    console.log('connected!');
    console.log(response);

    document.querySelector('.fb-login-button')[0].style.display = "none";
    document.querySelector('.login-spinner')[0].style.display = "none";
    document.querySelector('.login-continue')[0].style.display = "block";
    // response.authResponse.accessToken;
    // response.authResponse.userID;
    // TODO if on landing page:
    // TODO change login to 'continue to app'
    // TODO send id to server, set session cookie
    break;
  case 'not_authorized':
  case 'unkown':
  default:
    console.log('not connected');
    console.log(response);
    document.querySelector('.fb-login-button')[0].style.display = "block";
    document.querySelector('.login-spinner')[0].style.display = "none";
    document.querySelector('.login-continue')[0].style.display = "none";
    // TODO if on landing page:
    // TODO reset login button and spinner
    // TODO if not on landing page:
    //
    break;
  }
}

// ================== Logout Handlers ======================
function cleanLogout() {
  // TODO revoke session cookie
  // TODO Notify serve of logout
  window.location.replace("/");
}
function logout() {
  FB.logout(function(response) { cleanLogout(); });
}

// ================== Login Handlers ======================
function checkLoginState() {
  FB.getLoginStatus(function(response) { statusChangeCallback(response); });
}

// ================== Load SDK ======================
window.fbAsyncInit = function() {
  FB.init({
    appId : '367804320339821',
    cookie : true,
    xfbml : true,
    version : 'v2.11'
  });
  FB.AppEvents.logPageView();
  FB.getLoginStatus(function(response) { statusChangeCallback(response); });
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
