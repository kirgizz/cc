function login(){

}

function logout() {

}

function getArticle() {
  body = document.getElementById('body');
/*  section = document.createElement('section');
  section.setAttribute('class', 'common');

  name = document.createElement('div');
  name.setAttribute('class', 'name');

  author = document.createElement('div');
  author.setAttribute('class', 'author');

  data = document.createElement('div');
  data.setAttribute('class', 'data');

  article = document.createElement('acrticle');
  article.setAttribute('class', 'article');
  body.appendChild(section);
*/

  var request = new XMLHttpRequest();
 
  request.open('GET', 'api/getArticles', false);

  request.onload = function () {
    var data = JSON.parse(this.response); 

     if (request.status >= 200 && request.status < 400) {
        data.forEach(article => {

			  section = document.createElement('div');
  			section.setAttribute('class', 'common');

			  articleName = document.createElement('h1');
        articleName.setAttribute('class', 'articleName');
        articleName.textContent = article.Name;
        
			  author = document.createElement('h3');
        author.setAttribute('class', 'author');
        author.textContent = article.UserId;
				
				articleBody= document.createElement('article');
        articleBody.setAttribute('class', 'article');
        articleBody.textContent = article.Body;

				body.appendChild(section)
        section.appendChild(articleName);
        section.appendChild(author);
        section.appendChild(articleBody);

    }); 

  } else {
    const errorMessage = document.createElement('error');
    errorMessage.textContent = `Gah, it's not working!`;
    body.appendChild(errorMessage);
  }}
  request.send();


}

function addArticle() {

}

function updateArticle() {

}



function sendForm() {
		console.log("data")
/*    var http = new XMLHttpRequest();
    http.open("POST", "api/login", true);
    http.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    var params = document.getElementById('loginForm');
    http.send(params);
    alert(http.responseText);
*/
/*
    $.ajax({
        type:'post',//тип запроса: get,post либо head
        url:'api/login',//url адрес файла обработчика
        data:data,//параметры запроса
        processData: false,
        contentType: false,
        success:function(data){//возвращаемый результат от сервера
					alert("Success")
              },
        error:function(){
  				alert("Error")
        }
    });
*/
}


function call(form, url) {
  var msg   = $(form).serialize();
    $.ajax({
      type: 'POST',
      url: url,
      data: msg,
      success: function(data) {
      $('#results').html(data);
    },
    error:  function(xhr, str){
	    alert('Возникла ошибка: ' + xhr.responseCode);
    }
  });
}
