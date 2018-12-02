function createProfileForms() {
 
 profileForm = document.getElementByUd('profileForms')
 forms = ""


}

function getArticles() {
  body = document.getElementById('body');
  var request = new XMLHttpRequest();
 
  request.open('GET', 'api/getArticles', false);

  request.onload = function () {
    var data = JSON.parse(this.response); 
    if (request.status >= 200 && request.status < 400) {
//         $( "#body" ).append( "<p>i" + Test</p>" );
      data.forEach(article => {
        $("#body").append("<section class=\"article\"><div class=\"nickname\"><p>Author: " + article.Email + "<p></div><div class=\"article_name\"><p>Article name: " + article.Name + "<p></div><article class=\"article_body\"><p>" + article.Body + "<p></article></section>")

    }); 

  } else {
    const errorMessage = document.createElement('error');
    errorMessage.textContent = `Gah, it's not working!`;
    body.appendChild(errorMessage);
  }}
  request.send();
}
