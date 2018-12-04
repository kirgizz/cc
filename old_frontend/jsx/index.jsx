function HeadingItem (props) {
  const Heading = props.tagType
  return <Heading>{props.headingText}</Heading>
}

function DescriptionItem (props) {
  const Desc = props.tagType
  return <Desc>{props.descriptionText}</Desc>
}

function SectionItem (props) {
  return (<section className='item'>
    <HeadingItem
      headingText={props.headingText}
      tagType={props.headingTagType}
    />

    <DescriptionItem
      descriptionText={props.descriptionText}
      tagType={props.descriptionTagType}
    />
  </section>
	)
}

ReactDOM.render(
  <div className='wrapper'>
    <SectionItem
      headingText='HTML'
      headingTagType='h1'
      descriptionText='Transforms into the DOM'
      descriptionTagType='p'
    />

    <SectionItem
      headingText='CSS'
      headingTagType='h2'
      descriptionText='Style DOM-elements'
      descriptionTagType='quote'
    />

    <SectionItem
      headingText='JavaScript'
      headingTagType='h3'
      descriptionText='Adds interactivity to elements'
      descriptionTagType='div'
    />
  </div>,

  document.getElementById('app')
)
