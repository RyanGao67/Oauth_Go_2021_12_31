import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom";

import React from 'react';


class App extends React.Component{
    constructor(props) {
        super(props);
        this.state = {
            "access_token":"",
            "expires_in":"",
            "session_state":"",
            "token_type":""
        }
    }
    setStateValue = (k,v)=>{
        if(this.state[k]!==v){
            this.setState({[k]:v})
        }
    }
    onCheckStateClick = ()=>{
        console.log(this.state)
    }
    render (){
        return (
            <Router>
                <div>
                    <div className="App"><p>Implicit Grant Type</p></div>
                    <button onClick={this.onCheckStateClick }>Check state</button>
                    <nav>
                        <ul>
                            <li><Link to="/">home</Link></li>
                            <li><Link to="/login">Login</Link></li>
                            <li><Link to="/services">Services</Link></li>
                            <li><Link to="/logout">Logout</Link></li>
                        </ul>
                    </nav>

                    <Switch>
                        <Route path="/login"><Login/></Route>
                        <Route path="/callback"><Callback setStateValue={this.setStateValue}/></Route>
                        <Route path="/services"><Service accessToken={this.state.access_token}/></Route>
                        <Route path="/"><Home/></Route>
                    </Switch>
                </div>
            </Router>
        )
    }

}

function Home(){return <h2>Home</h2>;}
function Login(){
    window.location = "http://10.100.196.60:8080/auth/realms/learningApp/protocol/openid-connect/auth?client_id=implicitClient&response_type=token&redirect_uri=http://localhost:3000/callback&scope=getBillingService";
    return null;
}
function Callback(props){
    // Get access token
    console.log(window.location.hash)
    const hasStr = window.location.hash;
    const hashMap = hasStr.substr(1).split("&").reduce((accum, item) => {
      // add item to accumulator
        const kv = item.split("=")
        accum[kv[0]] = kv[1]
      // return accumulator
        return accum
    }, {})
    console.log(hashMap)
    const {setStateValue} = props;
    console.log(hashMap.access_token)
    setStateValue("access_token", hashMap.access_token)
    setStateValue("expires_in", hashMap.expires_in)
    setStateValue("session_state", hashMap.session_state)
    setStateValue("token_type", hashMap.token_type)
    return <div>Callback</div>
}

class Service extends React.Component{
    constructor(props) {
        super(props);
        this.state = {data:{}}
    }

    componentDidMount(){
        const {accessToken} = this.props

        console.log("tgao1")
        console.log(accessToken)
        console.log(this.props)
        const formData = new FormData()
        formData.append("access_token", accessToken)
        fetch("http://localhost:8081/billing/v1/services", {
            method: "POST",
            body: formData
        }).then(response => response.json()).then(data=>{
            console.log(data)
            this.setState({data})
        })

        const services = []
        services.push(<div key="1">billinga</div>)
        services.push(<div key="2">billinga</div>)
        services.push(<div key="3">billinga</div>)
    }
    render() {
        return <div>
            <h2>services</h2>
            <div>
                {JSON.stringify(this.state.data)}
            </div>
        </div>
    }
}



export default App;
