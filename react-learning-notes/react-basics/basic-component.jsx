// Named import: {} and exact element name, 'as' to rename.
// Default import: Without {}, from default export, rename by changing name.
import { OtherComponent } from 'directory/OtherComponent';

// Default export component.
// 'props' = optional parameter, object for passing properties.
export default function MyBasicComponent(props) {
  console.log(props.name);
  console.log(props.age);
  console.log(`my name is ${props.age}`)

  return (    // allows multi-line
    <>        {/* div alternative 'cause all elements need a top-container*/}
      <h1>Hello world!</h1>
      <p>My name is {props.name} and I'm {props.age} years old</p>
      <OtherComponent />     {/* use of imported components */}
    </>
  )
}

// ========================================================================
// Other component that imports MyBasicComponent
// isn't needed to write extension in from string
import MyBasicComponent from 'directory/MyBasicComponent';

export const MyFatherComponent = () => {
  return (
    <>
      <h1>Father Component</h1>
      {/* name and age value passed to MyBasicComponent props */}
      <MyBasicComponent name="Silas" age="27"/>
    </>
  )
}




