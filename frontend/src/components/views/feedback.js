import React from 'react';
import Select from 'react-select'
import CreatableSelect from 'react-select/lib/Creatable';

import "./styles/add.css";
import {Button, ControlLabel, FormControl, FormGroup} from "react-bootstrap";

import CKEditor from '@ckeditor/ckeditor5-react';
import ClassicEditor from '@ckeditor/ckeditor5-build-classic';


const editorConfiguration = {
    height: "500",
    toolbar: [ 'heading', '|', 'bold', 'italic', '|', 'undo', 'redo', ]
};


function feedback(props) {
    return (
        <div>
            <div className="info">
                <p>This is additional block</p>
                <p>This is additional block</p>
                <p>This is additional block</p>
                <p>This is additional block</p>
                <p>This is additional block</p>
            </div>
            <div className="block">

                <form onSubmit={props.handleSubmit}>

                    <FormGroup controlId="rubrics">
                        <div className="rubrics-select-form">
                            <p>Add rubrics to you publication</p>
                            <Select
                                value={props.state.rubrics}
                                onChange={props.handeChangeRubrics}
                                options={props.state.selectOptions}
                                isMulti={true}
                            />
                        </div>
                    </FormGroup>

                    <FormGroup controlId="title">
                        <div className="publication-title">
                            <ControlLabel>Title</ControlLabel>
                            <FormControl
                                autoFocus
                                value={props.state.title}
                                onChange={props.handleChangeTitle}
                            />

                        </div>
                    </FormGroup>


                    <FormGroup controlId="publicationBody">
                        <div className="publication-text-editor">
                            <p>Body</p>
                            <CKEditor
                                editor={ ClassicEditor }
                                onChange={props.handleEditorChange}
                                config={editorConfiguration}
                            />
                        </div>
                    </FormGroup>

                    <FormGroup controlId="tags">
                        <div className="tags-select-form">
                            <p>Add tags</p>
                            <CreatableSelect
                                isMulti
                                onChange={props.handleChangeTags}
                            />
                        </div>
                    </FormGroup>

                    <Button
                        block
                        bsSize="large"
                        type="submit"
                        //disabled={!props.validateForm()}
                        background-color="#4CAF50"
                    >

                        Sent
                    </Button>

                </form>
            </div>
        </div>
    )
}

export default feedback()
