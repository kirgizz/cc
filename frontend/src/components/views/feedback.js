import React from 'react';
import {Button, ControlLabel, FormControl, FormGroup} from "react-bootstrap";
import CKEditor from '@ckeditor/ckeditor5-react';
import ClassicEditor from '@ckeditor/ckeditor5-build-classic';


const editorConfiguration = {
    height: "500",
    toolbar: [ 'heading', '|', 'bold', 'italic', '|', 'undo', 'redo', ]
};

function handleChange( event ) {
    console.log(event)
}


function Feedback(props) {
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

                <form onSubmit={handleChange}>

                    <FormGroup controlId="title">
                        <div className="publication-title">
                            <ControlLabel>Title</ControlLabel>
                            <FormControl
                                autoFocus
                                //value={props.state.title}
                                onChange={handleChange}
                            />

                        </div>
                    </FormGroup>

                    <FormGroup controlId="email">
                        <div className="email">
                            <ControlLabel>Email</ControlLabel>
                            <FormControl
                                autoFocus
                                //value={props.state.title}
                                onChange={handleChange}
                            />

                        </div>
                    </FormGroup>

                    <FormGroup controlId="publicationBody">
                        <div className="publication-text-editor">
                            <CKEditor
                                editor={ ClassicEditor }
                                onChange={handleChange}
                                config={editorConfiguration}
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

export default Feedback
