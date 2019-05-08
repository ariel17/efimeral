import React, { Component } from 'react';
import { BrowserRouter, Route } from 'react-router-dom'
import Home from './containers/Home';
import About from './containers/About';
import Instance from './containers/Instance';

class App extends Component {

    render = () => {
        return (
          <BrowserRouter basename={process.env.PUBLIC_URL}>
              <Route exact path='/' component={Home} />
              <Route path='/about' component={About} />
              <Route path='/ubuntu' component={Instance} />
              <Route path='/arch' component={Instance} />
          </BrowserRouter>
        )
    }
}

export default App;
