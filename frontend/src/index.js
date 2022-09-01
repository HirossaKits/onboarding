import _ from 'lodash';

function component() {
  const element = document.createElement('div');
  element.innerHTML = _.json(['webpack', 'is'], 'working');
  return element;
}

document.body.appendChild(component());
