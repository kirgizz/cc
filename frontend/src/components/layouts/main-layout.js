import React from 'react';
import HeaderContainer from '../containers/header-container';
import FooterContainer from '../containers/footer-container';

export default function(props) {
 	 return (
 	 	<div>
 	 		<HeaderContainer />
 	       		{props.children}
 	 		<FooterContainer />
 	 	</div>
 	   );
}
