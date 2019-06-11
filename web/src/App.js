import React, { Component } from 'react';
import { HashRouter, Route } from 'react-router-dom'
import Home from './containers/Home';
import About from './containers/About';
import Instance from './containers/Instance';

class App extends Component {

    render = () => {
        return (
          <HashRouter basename={process.env.PUBLIC_URL}>
              <Route exact path='/' component={Home} />
              <Route path='/about' component={About} />
              <Route path='/ubuntu' render={
                (props) => <Instance distribution="ubuntu" tag="19.04" />
              } />
              <Route path='/arch' render={
                (props) => <Instance distro="arch" />
              } />
          </HashRouter>
        )
    }
}

export default App;
