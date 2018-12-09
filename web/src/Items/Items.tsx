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
    return this.props.items.map(item => {
      <li key={item.hash + '-' + Math.random()} className={cx('item', 'tooltip')} data-itemhash={item.hash}>
        <div className="icon">
          <ObservedImage className={cx('image', 'icon')} src={`https://www.bungie.net${props.icon}`} />
        </div>
      </li>
    });
  }
}

export default Items;
