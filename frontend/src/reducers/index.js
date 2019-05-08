import { combineReducers } from 'redux';

// Reducers
import articleReducer from './article-reducer';
import profileReducer from './profile-reducer';

// Combine Reducers
var reducers = combineReducers({
    articleState: articleReducer,
    profileState: profileReducer,
});

export default reducers;
