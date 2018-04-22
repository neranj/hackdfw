import React, {Component} from "react"
import axios from "axios"
import '../static/css/login.css'
import '../static/libs/bootstrap4/css/bootstrap.min.css'

class Login extends Component{

    handleClick(e){
        e.preventDefault();

        axios.get('http://localhost:8080/users/login')
            .then((data) => {
                console.log(data);
            })
            .catch((e) => {
                console.log(e);
            });

        return false;
    }

    render(){
        const self = this;

        return (
            <div id="login-container">
                <form className="form-signin" action="#" onSubmit={self.handleClick}>
                    <h2 className="form-signin-heading">Please sign in</h2>
                    <label htmlFor="inputEmail" className="sr-only">Email address</label>
                    <input type="email" id="inputEmail" className="form-control" placeholder="Email address" required
                           autoFocus />
                    <label htmlFor="inputPassword" className="sr-only">Password</label>
                    <input type="password" id="inputPassword" className="form-control" placeholder="Password" required />
                    <div className="checkbox">
                        <label>
                            <input type="checkbox" value="remember-me" /> Remember me
                        </label>
                    </div>
                    <button className="btn btn-lg btn-primary btn-block"
                            type="submit">Sign in</button>
                </form>

            </div>
        );
    }
}

export default Login;