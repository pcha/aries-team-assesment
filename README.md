# lapras
Lapras: Both a majestic water pokemon and the product list manager of our
🌈 dreams ✨

The purpose of this assignment is to mimic, as much as possible, a typical vertical
ticket on team Aires. Your ticket is FE-0001: As a user I can manage a list of products within
the Lapras UI.

The initial UI has already been completed; development instructions and acceptance criteria 
for the project are below. In order for this ticket to proceed to the next stage of the 
software development cycle on the road to production, all acceptance criteria must be met. 
At least one meaningful unit test must also be submitted. 

**The provided stack is just starting point, feel free to refactor or change code like a typical pull request.
If any big changes or additions are made, please document your thought process in the comments.**

## FE-0001
As a user I can manage a list of products within the Lapras UI

Background:
Our users need a new product manager to be able to quickly update and view product
details. This UI will be used in an internal tool accessed by users primarily 
on a Chrome desktop browser. They would like to be able to efficiently view & add products.

### Technical Notes
- Frontend application code should be imported into `./client/lapras-ui/app.tsx`
- Material UI is the UI component/styling library of choice for Lapras.
- Backend and database can be found in `./api/`

### Acceptance Criteria
- As a user I can see a list of products sorted by id, in a Table. For 
each product, I can see the id, name, and description
- As a user I can add a new product
- As a user, if a new product is successfully added, I can see an updated
product list
- As a member of the development team, I can run unit test(s) that meaningfully test the acceptance criteria

### Nice To Have / Bonus Points 
- Filter product table by search
- Adding auth in backend and/or frontend 

### Code Review Categories
Like a typical code review on our team, your code will be evaluated for the following categories: 

- Linting & Code Organization
- Testing Methodology
- Component Architecture
- Performance & Optimization
- Third Party Library Selection
- Scalability & Readability
- Enjoyable, Modern, and Nice User Experience & interface

### Requirements
- Go 1.18
- Node >= v16.14.2
- Yarn >=1.22.19
- Docker >= 20.10.14
- Docker compose >= v2.5.1

### How To Submit
- UI must be represented in a JSX (tsx) component imported and rendered into `./client/lapras-ui/App.tsx`
- Work on it as a normal repo, add your commits as usual **don't push your change into any remote service**
- Once you are done create a git bundle with: `git bundle create FirstnameLastname.git HEAD` and send the generated .git file to us
