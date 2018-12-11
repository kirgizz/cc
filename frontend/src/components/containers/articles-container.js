import React from 'react';
import { connect } from 'react-redux';
import Article from '../views/articles';
import * as articleApi from '../../api/articles-api';

const ArticleContainer = React.createClass({


  componentWillMount: function() {
    articleApi.getArticles();
    //store.dispatch(loadSearchLayout('userss', 'User Results'));
  },

  //componentDidMount: function() {
  //  articleApi.getArticles()
  //},

  render: function() {
    return (
      <Article articles={this.props.articles} />
    );
  }

});

const mapStateToProps = function(store) {
  return {
    articles: store.articleState.articles
  };
};

export default connect(mapStateToProps)(ArticleContainer);
