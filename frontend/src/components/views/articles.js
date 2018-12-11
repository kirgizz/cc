import React from 'react';
import { Link } from 'react-router';

// Using "Stateless Functional Components"
export default function(props) {
  return (
    <section className="Articles">


      {props.articles.map(article => {
             return (
              <div>
                  <div className="articleName">
                    <h1>{article.Name}</h1>
                  </div>

                  <div className="authorName">
                    <h2>{article.Email}</h2>
                  </div>

                  <div className="articleBody">
                    <p>{article.Body}</p>
                  </div>
              </div>
              );

        })}


    </section>
  );
}

/**{props.articles.map(article => {
       return (
        <div>
            <div className="articleName">
              <h1>{article.Name}</h1>
            </div>

            <div className="authorName">
              <h2>{article.Email}</h2>
            </div>

            <div className="articleBody">
              <p>{article.Body}</p>
            </div>
        </div>
        );

  })} */


//	type article struct{
//		Id			int
//		Email 		string
//		Name 		string
//		Rating 		int
//		Body 		string
//		View_count 	int
//	}
//      {props.articles.map(article => {
//return (
//    <div>
//      <div className="articleName">
//        <h1>{article.Name}</h1>
//      </div>

//      <div className="authorName">
//        <h2>{article.Email}</h2>
//      </div>

//      <div className="articleBody">
//        <p>{article.Body}</p>
//      </div>
//    </div>
//);

//})}