import React,{Component} from "React";
import axios from "axio";

import {Form,Icon,Card, Header, Input} from "semantic-ui-react"


let endpoint = 'http://localhost:9000'

class ToDoList extends Component{
    constructor(props){
        super(props)

        this.state = {
            task : "",
            items: [],
        }
    }


    ComponentDidMount(){
        this.getTask();
    }

    render(){
        return(
            <div className="row">
                <Header className="header" as="h2" color="yellow"> 
                To Do List
                </Header>

                 <form onSubmit={this.onSubmit}>
                    <input type="text" name="task" Onchange={this.Onchange} value={this.state.task} placeholder="create a task"></input>
                 </form>
            </div>
        )
    }
}