import React,{Component} from "react";
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

   Onchange = (event) => {
    this.SetState({[event.target.name]: event.target.value})
   }


getTask = ()=>{
    axios.get(endpoint+ '/api/task').then((re) => {
        if (res.date){
            this.setState({items:res.data.map((item) => {
                let color = 'yellow'
                let style = {
                    wordWrap : 'bread-world'
                }

                if (item.status){
                    color="green" 
                    style['textDecorationLine'] = 'line-through'

                }

                return (<Card key={item.id} color={color} fluid className='rough'>
                    <Card.Content>
                    <Card.Header textAling="left" ></Card.Header>

                    <Card.Meta textAling='right'>
                        <Icon name='check circle' color='blue' onClick={()=> this.updateTask(item,_id)} />
                            <span style={{paddingRight:10}}>
                                Undo
                            </span>
                     <Icon name = 'delete' color='red' onClick={() => this.deleteTask(item,_id)}/>

                     <span style={{paddingRight:10}}>Delete</span>
                    </Card.Meta>
                    </Card.Content>
                </Card> )
            })})
        } else {
            this.setState({items:[]})
        }
    })
}

updateTask = (id)=>{
    axios.put(endpoint+'/api/task'+id,{
        headers:{
            "Content-Type":"application/x-www-form-urlencoded",
        },
    } ).then((res)=>{
        console.log(res)
        this.getTask()
    })
}

undoTask = (id) => {
    axios.put(endpoint+'/api/undoTask'+id,{
        headers:{
            "Content-Type":"application/x-www-form-urlencoded",
        },
    } ).then((re) => {
        console.timeLog(re)
        this.getTask()
    })
}

deleteTask = (id) => {

        axios.delete(endpoint+"/api/DeleteTask", {header:{"Content-Type":"application/x-www-form-urlencoded"}}).then((re)=>{
            console.log(res)
            this.getTask()
        })
}

onSubmit = () => {
    let {task} = this.state

    if (task){
        axios.post(endpoint+"/api/task",{task}, {header:{"Content-Type":"application/x-www-form-urlencoded"}}).then((res)=>{
            this.getTask()
            this.setState({task:""})
        console.log(res)
        })

    }

}

    render(){
        return(
            <div className="row">
                <Header className="header" as="h2" color="yellow"> 
                To Do List
                </Header>
                   <div className="row">
                 <form onSubmit={this.onSubmit}>
                    <input type="text" name="task" Onchange={this.Onchange} value={this.state.task} placeholder="create a task"></input>
                {/* {<Button> Create Task</Button>} */}
                 </form>
            </div> 

            <div className="row">
                <Card.Group>{this.state.items}</Card.Group>
            </div>
            </div>
        )
    }
}