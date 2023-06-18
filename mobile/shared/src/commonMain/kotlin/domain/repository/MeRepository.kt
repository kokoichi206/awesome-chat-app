package domain.repository

import domain.model.User

interface MeRepository {

    suspend fun getMe(): User
}
