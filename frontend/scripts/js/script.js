function checkAuth() {
  if (document.cookie.indexOf("ssid") == 0) {
    var login = $('#login');
    var register = $('#register');
    login.text("logout");
    login.attr('href','/api/logout');
    register.text("Account");
    register.attr('href','/addArticle.html');
  } else {
    console.log("there is no cookie ssid")
  }
}

function getArticles() {
  body = document.getElementById('body');
  var request = new XMLHttpRequest();
 
  request.open('GET', 'api/getArticles', false);

  request.onload = function () {
    var data = JSON.parse(this.response); 
    var la = "test text"
    if (request.status >= 200 && request.status < 400) {
//         $( "#body" ).append( "<p>i" + Test</p>" );
      data.forEach(article => {
        $("#body").append("<section class=\"article\"><div class=\"nickname\"><p>Author: " + article.UserId + "<p></div><div class=\"article_name\"><p>Article name: " + article.Name + "<p></div><article class=\"article_body\"><p>" + article.Body + "<p></article></section>")

//			  section = document.createElement('section');
//  			section.setAttribute('class', 'article');
//
//        
//			  articleName = document.createElement('h1');
//        articleName.setAttribute('class', 'articleName');
//        articleName.textContent = article.Name;
//        
//			  author = document.createElement('h3');
//        author.setAttribute('class', 'author');
//        author.textContent = article.UserId;
//				
//				articleBody= document.createElement('article');
//        articleBody.setAttribute('class', 'article');
//        articleBody.textContent = article.Body;
//
//				body.appendChild(section)
//        section.appendChild(articleName);
//        section.appendChild(author);
//        section.appendChild(articleBody);

    }); 

  } else {
    const errorMessage = document.createElement('error');
    errorMessage.textContent = `Gah, it's not working!`;
    body.appendChild(errorMessage);
  }}
  request.send();
}

function call(form, url) {
  var f = $(form).serialize();
  console.log(f)
    $.ajax({
      type: 'POST',
      url: url,
      data: f,
    success:  function(xhr, str){
	    console.log("SUCCESS");
	    console.log(xhr.status);
	    console.log(xhr.responseText);
      if (url == "api/login") {
        window.location.replace("/index.html");
      }
    },
    error:  function(xhr, str){
	    console.log("ERROR");
	    console.log(xhr.status);
	    console.log(xhr.responseText);
    }
  });
}
