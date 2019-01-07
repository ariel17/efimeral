import React, { Component } from 'react';
import logo from './logo.svg';
import styled, { keyframes } from 'styled-components';

const Wrapper = styled.div`
    text-align: center;
`

const Header = styled.header`
    background-color: #282c34;
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    font-size: calc(10px + 2vmin);
    color: white;
`

const animation = keyframes`
    from {
      transform: rotate(0deg);
    }
    to {
      transform: rotate(360deg);
    }
`

const Logo = styled.img`
    animation: ${animation} infinite 20s linear;
    height: 40vmin;
`


class App extends Component {
  render() {
    return (
      <Wrapper>
        <Header>
          <Logo src={logo} alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </Header>
      </Wrapper>
    );
  }
}

export default App;
