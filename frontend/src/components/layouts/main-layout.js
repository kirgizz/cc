import React from 'react';
import Header from '../views/header'
import Footer from '../views/footer'
import HeaderContainer from '../containers/header-container';
import FooterContainer from '../containers/footer-container';
// Using "Stateless Functional Components"
export default function(props) {
 	 return (
 	 	<div>
 	 		<HeaderContainer />
 	       		{props.children}
 	 		<FooterContainer />
 	 	</div>
 	   );
}
