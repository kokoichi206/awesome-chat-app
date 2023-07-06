package di

import data.remote.RoomRemote
import data.remote.MeRemote
import data.repository.MeRepositoryImpl
import data.repository.RoomRepositoryImpl
import data.repository.SessionRepositoryImpl
import domain.repository.MeRepository
import domain.repository.RoomRepository
import domain.repository.SessionRepository
import domain.usecase.GetRoomMessagesUsecase
import domain.usecase.GetRoomUsersUsecase
import domain.usecase.PostMessagesUsecase
import org.koin.dsl.module
import util.provideDispatcher

private fun dataModule(platformModule: PlatformModule) = module {
    factory { MeRemote(get()) }
    factory { RoomRemote(get()) }

    factory<MeRepository> { MeRepositoryImpl(get()) }
    factory<RoomRepository> { RoomRepositoryImpl(get()) }
    factory<SessionRepository> { SessionRepositoryImpl(platformModule) }
}

private val domainModule = module {
    factory { GetRoomUsersUsecase(get()) }
    factory { GetRoomMessagesUsecase(get()) }
    factory { PostMessagesUsecase(get()) }
}

private val utilityModule = module {
    factory { provideDispatcher() }
}

fun sharedModules(platformModule: PlatformModule) = listOf(
    dataModule(platformModule),
    domainModule,
    utilityModule,
)

fun getSharedModules(platformModule: PlatformModule) =
    sharedModules(platformModule)
