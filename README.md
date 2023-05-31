# go-users-cleanarchitecture
Api en golang usando Clean Architecture



---Struct Diagram

@startuml

title Classes - Class Diagram

package routers{
  struct UserController {
    UserUseCase *usecases
  }
  object NewUserController{
    method()
  }
  object CreateUser{
    method()
  }
  UserController*..NewUserController
  UserController*..CreateUser
}

package domain{
  
  interface IUserRepository{}
  struct User{
  }
}



package usecases{
  struct UserUseCase{
    IUserRepository UserRepo
  }
  
  object NewUserUseCase{
    method(IUserRepository)
  }
  object Create{
    method()
  }
  UserUseCase *..NewUserUseCase
  UserUseCase *..Create
}

UserController-right->UserUseCase

package persistence{
  struct UserRepository{
    mongo.Collection collection
  }
  object NewUserRepository{
    method(mongo.Database)
  }
  object Insert{
    method(User)
  }
  UserRepository*..NewUserRepository
  UserRepository*..Insert
}

UserUseCase-right->IUserRepository
UserRepository-->User
IUserRepository<.. UserRepository

@enduml