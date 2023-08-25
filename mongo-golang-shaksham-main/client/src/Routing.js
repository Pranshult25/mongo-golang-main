import Header from "./Header";
import { BrowserRouter as Router, Redirect, Route, Switch } from "react-router-dom";
import RoutingSwitch from "./RoutingSwitch";
import PostFormModal from "./PostFormModal";
import AuthModal from "./AuthModal";
import { useContext, useEffect , useState } from "react";
import RedirectContext from "./RedirectContext";
import UserContext from "./UserContext"; // Import the UserContext
import ProfileModal from "./ProfileModal";


function Routing() {
  const { redirect, setRedirect } = useContext(RedirectContext);
  const user = useContext(UserContext); // Get user context
  const [showProfileModal, setShowProfileModal] = useState(false); // State to control profile modal

  useEffect(() => {
    if (redirect) {
      setRedirect(false);
    }
  }, [redirect]);

  return (
    <Router>
      {redirect && <Redirect to={redirect} />}
      {!redirect && (
        <>
          <Header setShowProfileModal={setShowProfileModal} /> {/* Pass setShowProfileModal to Header */}
          <RoutingSwitch />
          <PostFormModal />
          <AuthModal />
          
          {/* Show profile modal if user is logged in */}
          {user.username && (
            <Route
              path="/profile"
              render={() => (
                <ProfileModal onClose={() => setShowProfileModal(false)} />
              )}
            />
          )}
        </>
      )}
    </Router>
  );
}

export default Routing;
