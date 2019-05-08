import React from 'react';
import { connect } from 'react-redux';
import Article from '../views/articles';
import * as articleApi from '../../api/articles-api';


class ArticleContainer extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      articles: null,
      password: '',
      elemIsVisible: false,
      loginSuccess: true
    };
  }

  componentWillMount() {
    articleApi.getPublications().then(response => this.setState({ articles: response.data})).catch(console.log("cant get publications"))
  }

  render() {
    return (
      <Article />
    );
  }

};


const mapStateToProps = function(store) {
  return {
    articles: store.articleState.articles
  };
};

export default connect(mapStateToProps)(ArticleContainer);
