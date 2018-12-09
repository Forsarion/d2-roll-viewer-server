import React from 'react';
import cx from 'classnames';

import ObservedImage from './ObservedImage';

import './Items.css';

export interface Props {
  items: ItemProps[]
}

export interface ItemProps {
  hash: string
  icon: string
}

class Items extends React.Component<Props> {
  render() {    
    return <div id='items'>
      <ul>
          { 
            this.props.items.map(item => {
              return <li key={item.hash + '-' + Math.random()}>
                <div className='icon item tooltip' data-itemhash={item.hash}>
                  <ObservedImage className='image icon' src={`https://www.bungie.net${item.icon}`} />
                </div>
            </li>
          })
        }
      </ul>
    </div>
  }
}

export default Items;
