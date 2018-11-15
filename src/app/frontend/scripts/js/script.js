function login() {
  var xhr = new XMLHttpRequest();
 
  xhr.open('GET', '/login', false);
 
  xhr.send("la-la-la");
  
  if (xhr.status != 200) {
    alert( xhr.status + ': ' + xhr.statusText ); 
  } else {
    alert( xhr.responseText ); 
  }
	

}

function logout() {

}

function getArticle() {

}

function addArticle() {

}

function updateArticle() {

}


