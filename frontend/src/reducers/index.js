import { combineReducers } from 'redux';

// Reducers
import articleReducer from './article-reducer';

// Combine Reducers
var reducers = combineReducers({
    articleState: articleReducer,
});

export default reducers;
