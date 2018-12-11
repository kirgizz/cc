import React from 'react';
import { Link } from 'react-router';
import Header from '../views/header'
import Footer from '../views/footer'
// Using "Stateless Functional Components"
export default function(props) {
 	 return (
 	 	<div>
 	 		<Header />
 	       		{props.children}
 	 		<Footer />
 	 	</div>
 	   );
}
