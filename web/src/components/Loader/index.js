import React, { Component } from 'react';
import { css } from '@emotion/core';
import { BarLoader } from 'react-spinners';


const override = css`
    display: block;
    margin: 25% auto;
`;


class Loader extends Component {

    constructor(props) {
        super(props);
        this.state = {
            loading: props.loading
        }
    }

    render() {
        return <BarLoader
            css={override}
            sizeUnit={"px"}
            size={150}
            color={'#ffffff'}
            loading={this.state.loading} />
    }
}


export default Loader;
