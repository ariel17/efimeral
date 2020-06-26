import React, { Component } from 'react';
import { HashRouter, Route } from 'react-router-dom'
import About from './containers/About';
import Instance from './containers/Instance';

class App extends Component {

    render = () => {
        return (
          <HashRouter basename={process.env.PUBLIC_URL}>
              <Route path='/about' component={About} />
              <Route path='/' render={
                (props) => <Instance distribution="ubuntu" tag="20.04" />
              } />
          </HashRouter>
        )
    }
}

export default App;
