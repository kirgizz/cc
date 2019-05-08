import React from 'react';
import { EditorState} from 'draft-js';
import Add from '../views/add';
import {getRubrics} from "../../api/articles-api";

import {savePublication} from "../../api/articles-api";


class AddContainer extends React.Component {
    componentWillMount() {
        //TODO add processing exception
        getRubrics().then(response => this.setState({ selectOptions: response.data})).catch()

    }
    constructor(props) {
        super(props);
        this.state = {
            rubrics: [],
            title:'',
            body:'',
            selectOptions: [],
            tags: [],
            selectedValue: "",
            editorState: EditorState.createEmpty()
        };

        this.handeChangeRubrics = this.handeChangeRubrics.bind(this);
        this.handleChangeTitle = this.handleChangeTitle.bind(this);
        this.handleEditorChange = this.handleEditorChange.bind(this);
        this.handleChangeTags = this.handleChangeTags.bind(this);
    }

    handeChangeRubrics(event) {
        this.setState({
            rubrics: event
        })
    }

    handleChangeTitle = event => {
        this.setState({
            [event.target.id]: event.target.value
        });
    }

    handleEditorChange = (event, editor ) => {
        this.setState({
            body: editor.getData()
        })
    }

    handleChangeTags = (newValue, actionMeta) => {
        this.setState({
            tags: newValue,
        })
    };

    validateForm = () => this.state.rubrics.length > 0 && this.state.title.length > 0 && this.state.body.length > 0;

    handleSubmit = event => {
        event.preventDefault();
        savePublication({"rubrics":this.state.rubrics, "title": this.state.title, "body": this.state.body, "tags": this.state.tags}).then(
            function(result) {
                if (result.status === 200) {
                    console.log("Success")
                } else {
                    console.log("Error")
                }
            }.bind(this))
    }

    render() {
    return (
        <Add {...this}/>
    );
  }
}

export default AddContainer

//const mapStateToProps = function(store) {
//    return {
//        rubrics: store.articleState.articles
//    };
//};

//export default connect(mapStateToProps)(AddContainer);

