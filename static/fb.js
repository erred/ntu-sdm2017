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

// ======================== Cookie Get and Set
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

// ================================ FB API Callers
function updateAccountInfo(FB) {
  FB.api(
      '/me', {locale : 'en_US', fields : 'name, email, friends', limit : 100},
      function(response) {
        console.log(response);
        fetch('/account/update/', {
          method : 'post',
          headers : {
            'Accept' : 'application/json, text/plain, */*',
            'Content-Type' : 'application/json'
          },
          credentials : 'include',
          body :
              JSON.stringify({'name' : response.name, 'email' : response.email})

        })
            .then(function(res){console.log(res)});
      });
}

// function crawlFriends(FB) {
//   friends = [];
//   next = "/me/friends";
//   while (next != "") {
//     FB.api(next, function(response) {
//       response.data.forEach(function(el){friends.push(el)});
//       if ("next" in response.cursors) {
//         next = response.cursors.next;
//       } else {
//         next = "";
//       }
//     })
//   }
//   fetch('/account/friends/', {
//     method : 'post',
//     headers : {
//       'Accept' : 'application/json, text/plain, */*',
//       'Content-Type' : 'application/json'
//     },
//     credentials : 'include',
//     body : JSON.stringify(friends)
//   })
//       .then(function(res){return res.json()})
//       .then(function(res){console.log(res)});
// }

function setEmailName(FB) {
  FB.api('/me',
         {locale : 'en_US', fields : 'name, email, friends', limit : 100},
         function(response) {
           document.cookie = setCookie(100, "email", response.email);
           document.cookie = setCookie(100, "name", response.name);
         });
  if (getCookie("updateName") == "true") {
    updateAccountInfo(FB);
    setCookie(100, "updateName", "false");
  }
}

// ===================================== Main Dispatch

function statusChangeCallback(response) {
  switch (response.status) {
  case 'connected':
    setCookie(100, "id", response.authResponse.userID);
    setCookie(response.authResponse.expiresIn / (24 * 60 * 60), "token",
              response.authResponse.accessToken);
    if (window.location.pathname == "/") {
      setEmailName(FB);
      document.querySelector('.fb-login-button').style.display = "none";
      document.querySelector('.login-spinner').style.display = "none";
      document.querySelector('.login-continue').style.display = "block";
    }
    break;
  case 'not_authorized':
  case 'unkown':
  default:
    document.cookie = setCookie(-100, "id", "");
    document.cookie = setCookie(-100, "token", "");
    if (window.location.pathname == "/") {
      document.querySelector('.login-spinner').style.display = "none";
      document.querySelector('.fb-login-button').style.display = "block";
      document.querySelector('.login-continue').style.display = "none";
    } else {
      window.location.replace("/");
    }
  }
}

// ================== Login/Logout Handlers ======================
function logout() {
  FB.logout(function(response) { statusChangeCallback(response); });
}
function checkLoginState() {
  FB.getLoginStatus(function(response) { statusChangeCallback(response); });
}
