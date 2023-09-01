import BoardHeader from "./BoardHeader";
import PostForm from "./PostForm";
import PostsListingByCategory from "./PostListingByCategory";

function BoardCategory(props) {
  const {category} = props.match.params

  return (<div style={{minHeight:"100vh",backgroundColor:"black"}}>
    <BoardHeader />
    <PostForm />
    <PostsListingByCategory category={category}/>
  </div>);
}

export default BoardCategory;