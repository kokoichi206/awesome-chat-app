package data.repository

import data.remote.MeRemote
import data.remote.model.toUser
import domain.model.User
import domain.repository.MeRepository

class MeRepositoryImpl(
    private val remote: MeRemote,
) : MeRepository {

    override suspend fun getMe(): User {
        return remote.getMe().toUser()
    }
}