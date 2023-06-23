package di

import data.remote.RoomRemote
import data.remote.MeRemote
import data.repository.MeRepositoryImpl
import data.repository.RoomRepositoryImpl
import domain.repository.MeRepository
import domain.repository.RoomRepository
import domain.usecase.GetRoomMessagesUsecase
import domain.usecase.GetRoomUsersUsecase
import domain.usecase.PostMessagesUsecase
import org.koin.dsl.module
import util.provideDispatcher

private val dataModule = module {
    factory { MeRemote(get()) }
    factory { RoomRemote(get()) }

    factory<MeRepository> { MeRepositoryImpl(get()) }
    factory<RoomRepository> { RoomRepositoryImpl(get()) }
}

private val domainModule = module {
    factory { GetRoomUsersUsecase(get()) }
    factory { GetRoomMessagesUsecase(get()) }
    factory { PostMessagesUsecase(get()) }
}

private val utilityModule = module {
    factory { provideDispatcher() }
}

fun sharedModules() = listOf(
    dataModule,
    domainModule,
    utilityModule,
)

fun getSharedModules() =
    sharedModules()
