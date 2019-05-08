import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import Home from './containers/Home';
import About from './containers/About';

class App extends Component {

    render = () => {
        return (
          <BrowserRouter>
            <Switch>
              <Route exact path='/' render={
                (props) => <Home />
              } />
              <Route exact path='/about' render={
                (props) => <About />
              } />
            </Switch>
          </BrowserRouter>
        )
    }
}

export default App;
