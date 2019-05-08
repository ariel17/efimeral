import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import Home from './containers/Home';
import About from './containers/About';
import Instance from './containers/Instance';

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
              <Route exact path='/ubuntu' render={
                (props) => <Instance distro="ubuntu" />
              } />
              <Route exact path='/arch' render={
                (props) => <Instance distro="arch" />
              } />
            </Switch>
          </BrowserRouter>
        )
    }
}

export default App;
