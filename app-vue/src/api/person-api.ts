import client from '@/client/client'

class PersonApi {

  /**
   * Load all persons.
   */
  async loadAllPersons() {
    try {
      return await client.get('/api/v1/persons')
    } catch (err) {
      console.log('error while load all persons, ', err)
    }
  }

  /**
   * Remove person with selected ID.
   * @param id person identifier, uuid type.
   */
  async deletePerson(id: string) {
    try {
      return await client.delete(`api/v1/person/delete?id=${id}`)
    } catch (err) {
      console.log('error while delete person with id=', id, err)
    }
  }

  /**
   * Edit existing person
   * @param user information for backend:
   * {
   *     "id": "8ac045cd-a87b-472b-9f29-5a9f4b87", || uuid
   *     "firstName": "", || string
   *     "lastName": "", || string
   *     "age": 21 || number
   * }
   */
  async editPerson(user: object) {
    try {
      return await client.put(
        `/api/v1/person/update`,
        { id: user.id, firstName: user.firstName, lastName: user.lastName, age: user.age, login: user.login }
      )
    } catch (err) {
      console.log('error while edit person with id=', id, err)
    }
  }

  /**
   * Create new person.
   * @param user request to create new person
   * {
   *     "firstName": "",
   *     "lastName": "",
   *     "age": 21
   * }
   */
  async createPerson(user: object) {
    try {
      return await client.post(
        `/api/v1/person/create`,
        { firstName: user.firstName, lastName: user.lastName, age: user.age, login: user.login }
      )
    } catch (err) {
      console.log(`error while create new person, firstName=${user.firstName}, lastName=${user.lastName}, age=${user.age}`, err)
    }
  }
}

const personApi = new PersonApi()

export {
  personApi
}
