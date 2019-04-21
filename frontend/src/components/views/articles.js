import React from 'react';
import './styles/acrticles.css';
//import { animateScroll } from "react-scroll";

// Using "Stateless Functional Components"
export default function(props) {
  return (
    <section className="Articles">
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


      {props.articles.map(article => {
             return (
              <div>
                <div className="article">
                  <div className="articleName">
                    <h1>{article.Name}</h1>
                  </div>

                  <div className="authorName">
                    <h2>{article.Email}</h2>
                  </div>

                  <div className="articleBody">
                    <p>{article.Body}</p>
                  </div>
                  <div className="hr"></div>
                </div>
              </div>
              );

        })}


    </section>
  );
}
