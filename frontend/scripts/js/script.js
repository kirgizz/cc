function checkAuth() {
  if (document.cookie.indexOf("ssid") >= 0) {
    var login = $('#login');
    var register = $('#register');
    login.text("logout");
    login.attr('href','/api/logout');
    register.text("Account");
    register.attr('href','/addArticle.html');
  } else {
    //sent message to browser about wrong cookie
    console.log("there is no cookie ssid")
  }
}

function calculateRating(article, direction) {
    var articleRating = JSON.stringify({"article": article, "direction": direction});
    console.log(articleRating)
    $.ajax({
      type: 'POST',
      url: 'api/calculateRating',
      data: articleRating,
      async: true,
    success:  function(xhr, str){
      document.getElementById(article).innerHTML=xhr;
			//$("#this is my first article").text(xhr);
	    console.log("SUCCESS");
	    console.log(xhr);
    },
    error:  function(xhr, str){
	    console.log("ERROR");
	    console.log(xhr.status);
	    console.log(xhr.responseText);
    }
  });

}

function getArticles() {
  body = document.getElementById('body');
  var request = new XMLHttpRequest();
 
  request.open('GET', 'api/getArticles', false);

  request.onload = function () {
    var data = JSON.parse(this.response); 
    var la = "test text"
    console.log(data)
    if (request.status >= 200 && request.status < 400) {
//         $( "#body" ).append( "<p>i" + Test</p>" );
      data.forEach(article => {
        $("#body").append("<section class=\"article\">" + 
            "<div class=\"nickname\">" +
              "<p>Author: " + article.Email + "<p>" +
            "</div>" + 
            "<div class=\"article_name\">" + 
              "<p>Article name: " + article.Name + "<p>" +
            "</div>" +
            "<article class=\"article_body\">" +
              "<p>" + article.Body + "<p>" + 
              "<a href=\"article" + article.Id + "\"></a>" + 
            "</article>" + 
            "<div class=\"indicators\">" + 
              "<ul>" +
                "<li><button class=\"indicators-btn\" onclick=\"calculateRating('"+ article.Name+"',+1);\">UP</button>"+
                  "<span id=\""+article.Name+"\">" + article.Rating +  "</span><button class=\"indicators-btn\" onclick=\"calculateRating('"+ article.Name+"',-1);\">DOWN</button></li>" +
                "<li>C:0</li>" +
                "<li>V:"+ article.View_count + "</li>" +
              "</ul>" +
            "</div>" +
          "</section>");

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


function checkArticle() {
 var articleTextArea = $('#article_text_area');
 articleText = articleTextArea.val();
 html = $.parseHTML( articleText );
 //[""0""].nodeName
// console.log(html[0].data);
 var cutTag = 0
 $.each( html, function( key, value ) {
   if (html[key].nodeName == "CUT") {
	   if (html[key-1].textContent.length < 10) {
       console.log("text too littlet");
       return false;
     }
     return true
   }  
   //console.log( html[key].nodeName + ":" + html[key].textContent );
 });
 //console.log("==============")
 //console.log(html);
 return false
}

function addArticle(form) {
   var validateArticle = checkArticle()
   if (validateArticle = false) {
     console.log("error while validate article")
   } else {
     call(form, "api/addArticle")
   }

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
	    console.log(xhr);
    },
    error:  function(xhr, str){
	    console.log("ERROR");
	    console.log(xhr.status);
	    console.log(xhr.responseText);
    }
  });
}
