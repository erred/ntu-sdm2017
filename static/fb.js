// ================== StatusChange Handler ======================
function setCookie(exdays, key, value) {
  var userid = key + "=" + value;
  var d = new Date();
  d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
  var expires = "; expires=" + d.toUTCString();
  var path = "; path=/";
  document.cookie = userid + expires + path
}
function getCookie(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2)
    return parts.pop().split(";").shift();
  return ""
}

(function() {
  setCookie(100, 'id', '000000000');
  setCookie(100, 'name', 'A Name');
  setCookie(100, 'email', 'Example@example.com');

  if (window.location.pathname == "/") {
    document.querySelector('.fb-login-button').style.display = "none";
    document.querySelector('.login-spinner').style.display = "none";
    document.querySelector('.login-continue').style.display = "block";
  }

})();

function crawlFriends(FB) {
  friends = [];
  next = "/me/friends";
  while (next != "") {
    FB.api(next, function(response) {
      response.data.forEach(function(el){friends.push(el)});
      if ("next" in response.cursors) {
        next = response.cursors.next;
      } else {
        next = "";
      }
    })
  }
  fetch('/account/friends/', {
    method : 'post',
    headers : {
      'Accept' : 'application/json, text/plain, */*',
      'Content-Type' : 'application/json'
    },
    body : JSON.stringify(friends)
  })
      .then(function(res){return res.json()})
      .then(function(res){console.log(res)});
}

function statusChangeCallback(response) {
  switch (response.status) {
  case 'connected':
    console.log('connected!');
    console.log(response);

    // ================== Cookie ======================
    setCookie(100, "id", response.authResponse.userID);
    setCookie(response.authResponse.expiresIn / (24 * 60 * 60), "token",
              response.authResponse.accessToken);
    email = getCookie("email");
    if (email == "") {
      FB.api('/me',
             {locale : 'en_US', fields : 'name, email, friends', limit : 100},
             function(response) {
               document.cookie = setCookie(100, "email", response.email);
               document.cookie = setCookie(100, "name", response.name);
             });
    }

    // ================== UI Button ======================
    if (window.location.pathname == "/") {
      document.querySelector('.fb-login-button').style.display = "none";
      document.querySelector('.login-spinner').style.display = "none";
      document.querySelector('.login-continue').style.display = "block";
    }
    // ================== UI Button ======================
    if (window.location.pathname == "/account/") {
      crawlFriends(FB);
    }
    break;
  case 'not_authorized':
  case 'unkown':
  default:
    console.log('not connected');
    console.log(response);
    document.cookie = setCookie(-100, "id", "");
    document.cookie = setCookie(-100, "token", "");
    if (window.location.pathname == "/") {
      document.querySelector('.login-spinner').style.display = "none";
      document.querySelector('.fb-login-button').style.display = "block";
    } else {
      window.location.replace("/");
    }

    break;
  }
}

// ================== Logout Handlers ======================
function logout() {
  FB.logout(function(response) { statusChangeCallback(response); });
}
function logout() { window.location.replace("/"); }

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
