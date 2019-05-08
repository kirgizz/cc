import React from 'react';
import './styles/acrticles.css';

export default function(props) {
  return (
      <div>
          <div className="info">
              <p>I am cycle info block, a repeat alot</p>
              <p>I am cycle info block, a repeat alot</p>
              <p>I am cycle info block, a repeat alot</p>
              <p>I am cycle info block, a repeat alot</p>
              <p>I am cycle info block, a repeat alot</p>
              <p>I am cycle info block, a repeat alot</p>
              <p>I am cycle info block, a repeat alot</p>
              <p>I am cycle info block, a repeat alot</p>
          </div>
    <section className="block">




    </section>
          </div>
  );
}

//      {props.articles.map(article => {
//              return (
//               <div>
//                 <div className="article">
//                   <div className="articleName">
//                     <h1>{article.Name}</h1>
//                   </div>
//
//                   <div className="authorName">
//                     <h2>{article.Email}</h2>
//                   </div>
//
//                   <div className="articleBody">
//                     <p>{article.Body}</p>
//                   </div>
//                   <div className="hr"></div>
//                 </div>
//               </div>
//               );
//
//         })}
