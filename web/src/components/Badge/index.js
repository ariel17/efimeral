import React, { Component } from 'react';
import styled from 'styled-components';

const Img = styled.img`
    float: left;
`


class Badge extends Component {
    render() {
        return (<a href="https://github.com/ariel17/efimeral">
            <Img width="149" height="149" src="https://github.blog/wp-content/uploads/2008/12/forkme_left_white_ffffff.png?resize=149%2C149" class="attachment-full size-full" alt="Fork me on GitHub" data-recalc-dims="1" />
        </a>);

    }
}


export { Badge };
