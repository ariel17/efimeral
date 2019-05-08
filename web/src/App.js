import React, { Component } from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom'
import Home from './containers/Home';
import About from './containers/About';
import Instance from './containers/Instance';

class App extends Component {

    render = () => {
        return (
          <BrowserRouter basename={process.env.PUBLIC_URL}>
            <Switch>
              <Route exact path='/' render={
                (props) => <Home />
              } />
              <Route path='/about' render={
                (props) => <About />
              } />
              <Route path='/ubuntu' render={
                (props) => <Instance distro="ubuntu" />
              } />
              <Route path='/arch' render={
                (props) => <Instance distro="arch" />
              } />
            </Switch>
          </BrowserRouter>
        )
    }
}

export default App;
